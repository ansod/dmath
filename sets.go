package dmath

// Takes the supposed superset and subset
// and returns whether it is correct.
func IsSubset(superset, subset []int) bool {
    for _, val := range subset {
        if !Contains(val, superset) {
            return false
        }
    }
    return true
}

// Compares to sets(slices) and returns
// whether they are identical or not
func IsEqual(s1, s2 []int) bool {

    if len(s1) != len(s2) {return false}

    for _, val := range s1 {
        if !Contains(val, s2) {
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
                if !Contains(val, s) {
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
                if Contains(val, set2) && !Contains(val, s){
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

// Returns the complement between sets s1 and s2
func Complement(s1, s2 []int) []int {
    var s []int
    for _, val := range s1 {
        if !Contains(val, s2) {
            s = append(s, val)
        }
    }
    return s
}

// Returns whether set(slice) contains a
// certain number
func Contains(i int, set []int) bool {
    for _, v := range set {
        if v == i {
            return true
        }
    }
    return false
}
