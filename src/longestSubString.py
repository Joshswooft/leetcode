# Given a string s, find the length of the longest substring without repeating characters.

# Contraints
# 0 <= s.length <= 5 * 104
# s consists of English letters, digits, symbols and spaces.
def longestSubString(s: str)->int:
    # below code uses a sliding window approach

    # track the characters searched
    window_length = 0
    word_length = len(s)

    # track the longest non-repeating sub-string
    max_repeat = 0

    for i in range(word_length):
        # we move the window along by one char
        # if we find a repeating char we note its position and slide the window along
        # we carry on until we reach the end of the string
        # (or we could try and be smart and see how big the current string is and compare to length of string left to search)

        start = i
        # dict used to store word occurance
        w = {}
        counter = 0
        repeat = False
        for idx in range(start, word_length):
            key = s[idx]
            if w.get(key) != None:
                # found a repeating char in dict
                window_length = idx
                repeat = True
                if counter > max_repeat:
                    max_repeat = counter
                break
            else:
                w[key] = idx
            counter = counter+1

        # we looped through whole word with no repeats
        # print("sub string: ", s[start:])
        if repeat == False and max_repeat < len(s[start:]):
            max_repeat = counter
    return max_repeat
