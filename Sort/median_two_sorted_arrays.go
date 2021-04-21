func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    i, j, k := 0, 0, 0
    n, m := len(nums1), len(nums2)
    merged := make([]float64, n + m)
    
    // Use merge function from MergeSort
    for ; i < n && j < m; {
        if nums1[i] < nums2[j] {
            merged[k] = float64(nums1[i])
            i++
        } else {
            merged[k] = float64(nums2[j])
            j++
        }
        k++
    }
    for ; i < n; {
        merged[k] = float64(nums1[i])
        k++
        i++
    }
    for ; j < m; {
        merged[k] = float64(nums2[j])
        k++
        j++
    }
    
    // Calculate median
    if (n + m) % 2 != 0 {
        idx := int((n + m)/2)
        return merged[idx]
    } else {
        idx := int((n + m)/2)
        return (merged[idx] + merged[idx - 1])/2
    }
}