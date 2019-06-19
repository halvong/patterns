package singleton

import (
	"reflect"
	"testing"
)

func TestGetInstance(t *testing.T) {

	counter1 := GetInstance()

	if counter1 == nil {
		//Test 1 failed
		t.Error("A new connection object must have been made")
	}
	//expectedCounter := counter1
	t.Log("Type:",reflect.TypeOf(counter1))
	t.Log("Count:",counter1.count)

	//fmt.Println(expectedCounter)

	currentCount := counter1.AddOne()
	t.Log("Count:",counter1.count)
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	//---------------------------------
	counter2 := GetInstance()
	//if counter2 != expectedCounter
	if counter2 != counter1 {
		//Test 2 failed
		t.Error("Singleton instances must be different")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}

