# Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
# Difficulty: Hard
# The overall run time complexity should be O(log (m+n)).

# Constraints:

# nums1.length == m
# nums2.length == n
# 0 <= m <= 1000
# 0 <= n <= 1000
# 1 <= m + n <= 2000
# -106 <= nums1[i], nums2[i] <= 106
from pip import List


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        nums1.extend(nums2)
        nums1.sort()
        # "//" is floor division i.e. rounds number down
        mid = len(nums1) // 2
        return (nums1[mid] + nums1[~mid]) / 2
