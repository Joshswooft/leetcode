import unittest

from longestSubString import longestSubString

class TestLongestSubString(unittest.TestCase):


    def do_test(self, s:str, expectedLength:int)->bool:
        print("finding longest sub string: ", s)
        self.assertEqual(longestSubString(s), expectedLength)


    # Input: s = "abcabcbb"
    # Output: 3
    # Explanation: The answer is "abc", with the length of 3.
    def test_example_1(self):
        # abc
        self.do_test("abcabcbb", 3)
    
    # Input: s = "bbbbb"
    # Output: 1
    # Explanation: The answer is "b", with the length of 1.
    def test_example_2(self):
        # b
        self.do_test("bbbbb", 1)

    # Input: s = "pwwkew"
    # Output: 3
    # Explanation: The answer is "wke", with the length of 3.
    # Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
    def test_example_3(self):
        # wke
        self.do_test("pwwkew", 3)

    def text_example_4(self):
        self.do_test("a", 1)

    def test_no_repeats(self):
        self.do_test("au", 2)

    def test_no_repeats_in_sub_word(self):
        # ab
        self.do_test("aab", 2)

    def test_example_dvdf(self):
        # vdf
        # problem i have: it searches 'dvd' then will move on to 'df'
        # how it should be: 'dv', 'vdf'
        # how it is: 'dv', 'df'
        self.do_test("dvdf", 3)

if __name__ == '__main__':
    unittest.main()