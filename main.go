package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	splits    []time.Duration
	isRunning bool
	timeStamp int64
}

// Start Запускает и выключает секундомер
func (s *Stopwatch) Start() {
	t := time.Now()

	// Если секундомер выключен, то сохраняем временной штамп начала замера в наносекундах
	if !s.isRunning {
		s.timeStamp = t.UnixNano()
	}
	// Если секундомер уже запущен, то обнуляем наши замеры
	if s.isRunning {
		s.splits = []time.Duration{}
	}
	// Выключаем секундомер (Если уже выключен, то, наоборот, включится)
	s.isRunning = !s.isRunning
}

// SaveSplit Сохранение промежуточного замера секундомера
func (s *Stopwatch) SaveSplit() {
	t := time.Now()
	// Приводим к типу time.Duration результат вычитания текущего времени и времени запуска секундомера
	s.splits = append(s.splits, time.Duration(t.UnixNano()-s.timeStamp))
}

// GetResults Получаем сохраненные замеры в формате time.Duration
func (s *Stopwatch) GetResults() []time.Duration {
	return s.splits
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())

	// Выключение
	sw.Start()
	// Включение
	sw.Start()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()
	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
