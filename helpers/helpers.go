package helpers

import "github.com/barathsurya2004/sproutsland/objects"

func RemoveObjects(arr []objects.Object, i int) []objects.Object {
	copy(arr[i:], arr[i+1:])
	return arr[:len(arr)-1]
}
