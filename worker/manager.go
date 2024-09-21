package main

type Manager struct {
	ID string
	addr string

}

var managers []Manager

func isContainInManagerList(_id string) int {
	for index, manager := range managers {
		if manager.ID == _id {
			return index
		}
	}
	return -1
}

func addToManagerList(manager Manager) {
	if isContainInManagerList(manager.ID) != -1 {
		manager.append(managers, manager)
	}
}

func removeFromManagerList(_id string) {
	index := isContainInManagerList(_id)
	if index != -1 {
		managers = append(managers[:index], managers[index + 1:]...)
	}
} 