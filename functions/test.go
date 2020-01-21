package functions

// TestStruct ...
type TestStruct struct {
	Name string `json:"name,omitempty"`
}

// TestNil ...
func TestNil(shouldReturnNil bool) (testStructSlice []TestStruct) {

	if shouldReturnNil {
		return nil
	}
	return append(testStructSlice, TestStruct{
		Name: "Kewin",
	})
}
