# Given a string s, find the length of the longest substring without repeating characters.
# Difficulty: Medium

# Contraints
# 0 <= s.length <= 5 * 104
# s consists of English letters, digits, symbols and spaces.
# below code uses a sliding window approach
# we move the end pointer along filling up the dictionary with char->index mappings and update the max sub-string length 
# if more than current max
# if we encounter a duplicate then we update the start position to be the index of the encountered duplicate +1
# if the start pointer is infront of the end pointer then we increase the end pointer
def longestSubString(s: str)->int:
    word_length = len(s)

    # track the longest non-repeating sub-string
    n = 0

    # pointers for the sliding window
    start = 0
    end = 0
    # dict used to store word occurance - used to find duplicates
    w = {}

    while end < word_length:

        char = s[end]
        if char not in w or start > w[char]:
            w[char] = end
            n = max(n, end - start + 1)
            end = end + 1
        else:
            start = w[char] + 1

    return n
