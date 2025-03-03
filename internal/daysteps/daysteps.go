package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("conversion error: %w\ndata: %s", err, datastring)
	}

	ds.Steps, err = strconv.Atoi(parts[0])
	ds.Duration, err = time.ParseDuration(parts[1])
	return
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() string {
	distance := spentenergy.Distance(ds.Steps)
	kal := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал. \n",
		ds.Steps, distance, kal,
	)
}
