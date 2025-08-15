package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    // Solution 1
    var fav []int
    fav = []int{2,6,9}
    return fav
    /* Solution 2
    fav := []int{2,6,9}
    return fav
    */
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    var x int
    y := len(slice)
    if index > y || index < 0{
    	return -1
    } else{
        x = slice[index]
        return x
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    var a []int
    y := len(slice)
    if index >= y || index < 0{
        a = append(slice, value)
        return a
    } else {
        slice[index] = value
        return slice
    }
    /*
     Alternate solution:
    y := len(slice)
    if index >= y || index < 0{
        slice = append(slice, value)
        return slice
    } else {
        slice[index] = value
        return slice
    }
    */
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    prepend := append(values, slice...) //Need to use the unpack operator "..." after slice for the append command to work correctly.
    return prepend
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    var y int
    y = len(slice)
    if index > y || index < 0{
        return slice
    } else {
        preSlice := slice[:index]
        postSlice := slice[index+1:]
        newSlice:= append(preSlice, postSlice...)
        return newSlice
    }
}
