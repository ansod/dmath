package dmath

// Takes the supposed superset and subset
// and returns whether it is correct.
func IsSubset(superset, subset []int) bool {
    for _, val := range subset {
        if !sliceContains(val, superset) {
            return false
        }
    }
    return true
}

// Returns the union of an arbitrary
// number of sets(slices).
func Union(sets ...[]int) []int {
    var s []int

    if len(sets) > 0 {
        for _, set := range sets {
            for _, val := range set {
                if !sliceContains(val, s) {
                    s = append(s, val)
                }
            }
        }
    }
    return s
}

// Returns the intersection of an arbitrary
// number of sets(slices).
func Intersection(sets ...[]int) []int {
    var s []int

    for _, set := range sets {
        for _, val := range set {
            occurrences := 0
            for _, set2 := range sets {
                if sliceContains(val, set2) && !sliceContains(val, s){
                    occurrences++
                }
            }
            if occurrences == len(sets) {
                s = append(s, val)
            }
        }
    }
    return s
}

// Returns whether set(slice) contains a
// certain number
func sliceContains(i int, set []int) bool {
    for _, v := range set {
        if v == i {
            return true
        }
    }
    return false
}
