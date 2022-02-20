# Given a string s, find the length of the longest substring without repeating characters.

# Contraints
# 0 <= s.length <= 5 * 104
# s consists of English letters, digits, symbols and spaces.
def longestSubString(s: str)->int:

# how to find the longest sub string which doesnt repeat?
# i.e. abcabcbb
# a -> b -> c
# a -> b -> c -> a
# a -> b -> c -> a -> b
# a -> b -> c -> a -> b -> c (repeat detected)

# use a sliding window approach
    # track the characters searched
    window_length = 0
    word_length = len(s)
    max_repeat = 0
    for i in range(word_length):
        # we move the window along by one char
        # if we find a repeating char we note its position and slide the window along
        # we carry on until we reach the end of the string
        # (or we could try and be smart and see how big the current string is and compare to length of string left to search)

        start = window_length
        # dict used to store word occurance
        w = {}
        counter = 0
        for idx in range(start, word_length):
            key = s[idx]
            if w.get(key) != None:
                # found a repeating char
                window_length = idx
                if counter > max_repeat:
                    max_repeat = counter
                break
            else:
                w[key] = idx
            counter = counter+1

    return max_repeat
