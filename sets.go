package dmath


func Union(sets ...[]int) []int {
    var s1 []int

    if len(sets) > 0 {
        for _, set := range sets {
            for _, val := range set {
                if !sliceContains(val, s1) {
                    s1 = append(s1, val)
                }
            }
        }
    }
    return s1
}


func sliceContains(i int, set []int) bool {
    for _,v := range set {
        if v == i {
            return true
        }
    }
    return false
}
