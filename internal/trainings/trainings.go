package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("conversion error: %w\ndata: %s", err, datastring)
	}

	t.Steps, err = strconv.Atoi(parts[0])

	t.TrainingType = parts[1]

	t.Duration, err = time.ParseDuration(parts[2])

	return
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() string {
	distance := spentenergy.Distance(t.Steps)
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var kal float64
	switch t.TrainingType {
	case "Ходьба":
		kal = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		kal = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	default:
		return fmt.Sprintf("Неизвестный тип тренировки: %s", t.TrainingType)
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, kal,
	)
}
