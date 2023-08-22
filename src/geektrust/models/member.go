package expense

import "sort"

type Member struct {
	name string
	dues map[string]int
}

func NewMember(name string) *Member {
	return &Member{
		name: name,
		dues: make(map[string]int),
	}
}

func (m *Member) AddDue(member string, amount int) {
	m.dues[member] += amount
}

func (m *Member) ClearDue(member string, amount int) bool {
	if m.dues[member] >= amount {
		m.dues[member] -= amount
		return true
	}
	return false
}

func (m *Member) GetDues() []MemberDue {
	dues := make([]MemberDue, 0)
	for member, amount := range m.dues {
		dues = append(dues, MemberDue{member, amount})
	}
	sort.Slice(dues, func(i, j int) bool {
		if dues[i].Amount != dues[j].Amount {
			return dues[i].Amount > dues[j].Amount
		}
		return dues[i].Member < dues[j].Member
	})
	return dues
}

type MemberDue struct {
	Member string
	Amount int
}
