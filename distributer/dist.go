package distributer

import (
	"Qubecinema/challenge/utils"
	"fmt"
	"strings"
)

type Distributor struct {
	Name        string
	Parent      string
	Permissions map[string]bool
}

var DistMap = make(map[string]Distributor)

func AddDistributor(Name, Parent string, Include, Exclude []string) error {
	if _, exists := DistMap[Name]; exists {
		return fmt.Errorf("distributor %s already exists", Name)
	}
	newDist := Distributor{
		Name:        Name,
		Parent:      Parent,
		Permissions: make(map[string]bool),
	}
	if Parent != "" {
		parentDist, exists := DistMap[Parent]
		if !exists {
			return fmt.Errorf("parent distributor %s not found", Parent)
		}
		for _, region := range Include {
			if !parentDist.Permissions[region] {
				return fmt.Errorf("parent distributor %s does not have access to %s", Parent, region)
			}
		}
	}
	for _, region := range Include {
		newDist.Permissions[region] = true
		for area := range utils.AreaMap {
			if strings.HasSuffix(area, region) {
				newDist.Permissions[area] = true
			}
		}
	}
	for _, region := range Exclude {
		for area := range newDist.Permissions {
			if strings.HasSuffix(area, region) {
				delete(newDist.Permissions, area)
			}
		}
	}
	DistMap[Name] = newDist
	return nil
}

func CheckPermission(distributorName, region string) (bool, error) {
	dist, exists := DistMap[distributorName]
	if !exists {
		return false, fmt.Errorf("distributor %s not found", distributorName)
	}

	if allowed, found := dist.Permissions[region]; found {
		return allowed, nil
	}

	return false, nil
}
