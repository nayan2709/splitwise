package expense

import (
	"fmt"
	constants "geektrust/constants"
	model "geektrust/models"
	"strconv"
)

type expenseHandler struct {
	House *model.House
}

func NewExpenseHandler() IExpenseHandler {
	return &expenseHandler{
		House: model.NewHouse(),
	}
}

// IExpenseHandler is an interface for ExpenseHandler
type IExpenseHandler interface {
	ExecuteOperation(args []string)
}

func (handler *expenseHandler) ExecuteOperation(args []string) {
	var output string
	switch args[0] {
	case constants.MoveIn:
		output = handler.House.MoveIn(args[1])
	case constants.MoveOut:
		output = handler.House.MoveOut(args[1])
	case constants.Spend:
		{
			amount, err := strconv.Atoi(args[1])
			if err != nil {
				output = constants.InvalidAmount
				break
			}
			output = handler.House.Spend(amount, args[2], args[3:])
		}
	case constants.Dues:
		output = handler.House.Dues(args[1])
	case constants.ClearDue:
		{
			amount, err := strconv.Atoi(args[3])
			if err != nil {
				output = constants.InvalidAmount
				break
			}
			output = handler.House.ClearDue(args[1], args[2], amount)
		}
	default:
		output = constants.InvalidOperation
	}
	fmt.Println(output)
}
