# Given a string s, find the length of the longest substring without repeating characters.

# Contraints
# 0 <= s.length <= 5 * 104
# s consists of English letters, digits, symbols and spaces.
# below code uses a sliding window approach
# we move the window along by one char
# if we find a repeating char we note its position and slide the window along
# we carry on until we reach the end of the string
def longestSubString(s: str)->int:
    word_length = len(s)

    # track the longest non-repeating sub-string
    max_repeat = 0

    # pointers for the sliding window
    start = 0

    for _ in range(word_length):

        # early exit
        if start + max_repeat >= word_length:
            return max_repeat

        # dict used to store word occurance - used to find duplicates
        w = {}
        counter = 0
        for idx in range(start, word_length):
            key = s[idx]
            if w.get(key) != None:
                # found a repeating char in dict
                # we move the start position to the next character along from the duplicate index
                start = w.get(key) + 1
                if counter > max_repeat:
                    max_repeat = counter
                break
            else:
                w[key] = idx
            counter = counter+1
            if counter > max_repeat:
                max_repeat = counter

    return max_repeat
