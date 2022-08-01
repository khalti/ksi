package util

// Checks if object is nil
func Nil(a any) bool {
	return a == nil
}

// Checks if object is not nil
func NotNil(a any) bool {
	return !Nil(a)
}