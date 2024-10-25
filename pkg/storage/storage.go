package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"go.uber.org/zap"
)

type Storage struct {
	data  map[string]string
	lists map[string][]int
}

func NewStorage() *Storage {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("NewStorage")
	return &Storage{data: make(map[string]string), lists: make(map[string][]int)}
}

func (s *Storage) Set(key, value string) bool {
	s.data[key] = value
	return true
}

func (s *Storage) GetList(key string) []int {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	val, err := s.lists[key]
	if !err {
		logger.Info("Error in get")
		return nil
	}
	return val
}

func (s *Storage) Get(key string) *string {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	val, err := s.data[key]
	if !err {
		logger.Info("Error in get")
		return nil
	}
	return &val
}

func (s *Storage) GetKind(key string) string {
	if _, err := strconv.Atoi(s.data[key]); err != nil {
		return "S"
	}
	return "D"
}

func (s *Storage) LPUSH(list string, add []int) int {
	s.lists[list] = append(add, s.lists[list]...)
	return len(s.lists[list])
}

func (s *Storage) RPUSH(list string, add []int) int {
	s.lists[list] = append(s.lists[list], add...)
	return len(s.lists[list])
}

func (s *Storage) RADDTOSET(list string, add []int) {
	arr := s.lists[list]
	for i := 0; i < len(add); i++ {
		flag := true
		for j := 0; j < len(arr); j++ {
			if arr[j] == add[i] {
				flag = false
				break
			}
		}
		if flag {
			arr = append(arr, add[i])
		}
	}
	s.lists[list] = arr

}

func (s *Storage) LPOP(key string, bands []int) []int {
	var ret []int
	band := make([]int, len(bands))
	copy(band, bands)

	if len(band) > 1 {
		s.checkMinus(band, key)
		if s.checkRange(band, key) {
			return []int{len(s.lists[key])}
		}
		ret := make([]int, len(s.lists[key][band[0]:band[1]+1]))
		copy(ret, s.lists[key][band[0]:band[1]+1])
		s.lists[key] = append(s.lists[key][:band[0]], s.lists[key][band[1]+1:]...)

		return ret
	}
	if band[0] < 0 {
		if s.checkRange(band, key) {
			return []int{len(s.lists[key])}
		}
		band[0] = len(s.lists[key]) + band[0]
	}
	ret = make([]int, len(s.lists[key][:band[0]]))

	copy(ret, s.lists[key][:band[0]])
	s.lists[key] = s.lists[key][band[0]:]
	return ret
}

func (s *Storage) checkMinus(band []int, key string) {
	if band[0] < 0 && band[1] < 0 {
		band[0] = len(s.lists[key]) + band[0]
		band[1] = len(s.lists[key]) + band[1]

	} else if band[0] < 0 && band[1] >= 0 {
		band[0] = len(s.lists[key]) + band[0]

	} else if band[0] >= 0 && band[1] < 0 {
		band[1] = len(s.lists[key]) + band[1]
	}

}

func (s *Storage) checkRange(band []int, key string) bool {
	if len(band) == 0 || band[1] >= len(s.lists[key]) ||
		band[0] > band[1] || band[0] < 0 {
		return true
	}
	if len(band) == 1 &&
		(band[0] > len(s.lists[key]) || -band[0] > len(s.lists[key])) {
		return true
	}
	return false
}

func (s *Storage) RPOP(key string, bands []int) []int {
	band := make([]int, len(bands))
	copy(band, bands)
	var ret []int
	if len(band) > 1 {
		s.checkMinus(band, key)
		if s.checkRange(band, key) {
			return []int{len(s.lists[key])}
		}
		ret := make([]int, len(s.lists[key][band[0]:band[1]+1]))
		copy(ret, s.lists[key][band[0]:band[1]+1])
		s.lists[key] = append(s.lists[key][:band[0]], s.lists[key][band[1]+1:]...)
		return ret
	}
	if band[0] < 0 {
		if s.checkRange(band, key) {
			return []int{len(s.lists[key])}
		}
		band[0] = len(s.lists[key]) + band[0]
	}
	band[0] = len(s.lists[key]) - band[0]
	ret = make([]int, len(s.lists[key][band[0]:]))
	copy(ret, s.lists[key][band[0]:])
	s.lists[key] = s.lists[key][:band[0]]
	return ret

}

func (s *Storage) LSET(key string, index int, element int) string {
	if _, exists := s.lists[key]; !exists {
		return "wrong key"
	}
	if index < 0 || index >= len(s.lists[key]) {
		return "Index out of range"
	}
	s.lists[key][index] = element
	return "OK"

}

func (s *Storage) LGET(key string, index int) (int, bool) {
	if index < 0 || index >= len(s.lists[key]) {
		return 0, true
	}
	ret := s.lists[key][index]
	return ret , false

}

func (s *Storage) Serialization() {
	file, err := os.OpenFile("../pkg/storage/lists.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("error in openning file ", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(s.lists)
	if err != nil {
		fmt.Println("Error encoding in json ", err)
		return
	}
	fmt.Println("Data successfully serialized")

}

func (s *Storage) Deseserialization() {
	file, err := os.OpenFile("../pkg/storage/lists.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("error in openning file ", err)
		return

	}

	decoder := json.NewDecoder(file)
	decoder.Decode(&s.lists)

}
