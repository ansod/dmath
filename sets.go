package dmath


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

func sliceContains(i int, set []int) bool {
    for _, v := range set {
        if v == i {
            return true
        }
    }
    return false
}
