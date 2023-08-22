package expense

import (
	"fmt"
	constants "geektrust/constants"
	"math"
)

type House struct {
	members map[string]*Member
}

func NewHouse() *House {
	return &House{
		members: make(map[string]*Member),
	}
}

func (h *House) MoveIn(name string) string {
	if len(h.members) >= 3 {
		return constants.Houseful
	}
	h.members[name] = NewMember(name)
	for member := range h.members {
		if member != name {
			h.members[member].AddDue(name, 0)
		}
	}
	return constants.Success
}

func (h *House) Spend(amount int, spentBy string, spentFor []string) string {

	if _, ok := h.members[spentBy]; !ok || len(h.members) < 2 {
		return constants.MemberNotFound
	}
	for _, member := range spentFor {
		if _, ok := h.members[member]; !ok {
			return constants.MemberNotFound
		}
	}

	totalMembers := len(spentFor) + 1
	perPersonShare := amount / totalMembers

	for _, member := range spentFor {
		h.members[member].AddDue(spentBy, perPersonShare)
	}
	h.OptimiseDues()
	return constants.Success
}

func (h *House) Dues(memberName string) string {
	member, ok := h.members[memberName]
	if !ok {
		return constants.MemberNotFound
	}

	dues := member.GetDues()
	result := ""
	for i, due := range dues {
		if i < len(dues)-1 {
			result += fmt.Sprintf("%s %d\n", due.Member, due.Amount)
		} else {
			result += fmt.Sprintf("%s %d", due.Member, due.Amount)
		}
	}
	return result
}

func (h *House) ClearDue(memberWhoOwes, memberWhoLent string, amount int) string {
	member, ok1 := h.members[memberWhoOwes]
	_, ok2 := h.members[memberWhoLent]
	if !ok1 || !ok2 {
		return constants.MemberNotFound
	}

	if !member.ClearDue(memberWhoLent, amount) {
		return constants.InvalidAmount
	}

	return fmt.Sprintf("%d", member.dues[memberWhoLent])
}

func (h *House) MoveOut(name string) string {
	member, ok := h.members[name]
	if !ok {
		return constants.MemberNotFound
	}

	for otherMember, dueAmount := range member.dues {
		if dueAmount != 0 || h.members[otherMember].dues[name] != 0 {
			return constants.Failure
		}
	}

	delete(h.members, name)
	return constants.Success
}

func (h *House) OptimiseDues() {
	dueAmounts := make(map[string]int)
	for _, member := range h.members {
		dues := member.GetDues()
		for _, due := range dues {
			if due.Amount == 0 {
				continue
			}
			dueAmounts[member.name] += due.Amount
			dueAmounts[due.Member] -= due.Amount
			h.ClearDue(member.name, due.Member, due.Amount)
		}
	}
	h.minCashFlowRec(dueAmounts)
}

func (h *House) minCashFlowRec(amounts map[string]int) {
	var mxCredit, mxDebit string

	for person := range amounts {
		if amounts[person] > amounts[mxCredit] {
			mxCredit = person
		}
		if amounts[person] < amounts[mxDebit] {
			mxDebit = person
		}
	}

	if amounts[mxCredit] == 0 && amounts[mxDebit] == 0 {
		return
	}

	minAmount := int(math.Min(float64(-amounts[mxDebit]), float64(amounts[mxCredit])))
	amounts[mxCredit] -= minAmount
	amounts[mxDebit] += minAmount

	h.members[mxCredit].AddDue(mxDebit, minAmount)

	h.minCashFlowRec(amounts)
}
