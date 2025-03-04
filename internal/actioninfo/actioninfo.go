package actioninfo

import (
	"fmt"
)

// DataParser — интерфейс, описывающий методы Parse и ActionInfo.
type DataParser interface {
	Parse(string) error
	ActionInfo() string
}

// Info — функция для обработки данных о тренировках или прогулках.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Попытка распарсить строку.
		err := dp.Parse(data)
		if err != nil {
			// Если ошибка, переходим к следующей итерации.
			fmt.Printf("conversion error: %w\n\n", err)
			continue
		}

		// Формируем и выводим информацию об активности.
		fmt.Println(dp.ActionInfo())
	}
}
