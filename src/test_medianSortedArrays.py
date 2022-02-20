import unittest

from medianSortedArrays import Solution


class TestMedianSortedArrays(unittest.TestCase):

    sol = Solution()

    # note: constraint is one of the arrays has to have number
    def test_empty_arrays(self):
        res = self.sol.findMedianSortedArrays([],[1])
        self.assertEqual(res, 1)

        res2 = self.sol.findMedianSortedArrays([2], [])
        self.assertEqual(res2, 2)

    def test_example_1(self):
        res = self.sol.findMedianSortedArrays([1, 3], [2])
        # merged array = [1,2,3], median = 2
        self.assertEqual(res, 2)

    def test_example_2(self):
        res = self.sol.findMedianSortedArrays([1,2], [3,4])
        self.assertEqual(res, 2.5)

    # one of the rules is that nums can be negative
    def test_neg_example(self):
        res = self.sol.findMedianSortedArrays([-3,-2,-1], [-1, 1, 2])
        self.assertEqual(res, -1)

if __name__ == '__main__':
    unittest.main()