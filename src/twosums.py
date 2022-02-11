class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        dict = {}
        # always use a hash map!
        for index, number in enumerate(nums):
            if target - number in dict:
                # assumes there to be one exact solution so we can return here
                return [dict.get(target-number),index]
            dict[number] = index
        return