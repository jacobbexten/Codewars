def get_count(sentence):
    count = 0
    vowels = "aeiou"
    for i in sentence:
        if i in vowels:
            count = count + 1
    return count
