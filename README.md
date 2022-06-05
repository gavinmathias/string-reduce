# Zuma (one dimension candy crash game)

## Given a string representing a list of single digit, remove any group of same digit if there are at least 3 consecutive occurrence. Repeat the same treatment to the string after removal until there is no more group of same digit with 3 or more consecutive occurrence

Example 1:
     input: 112233444321
     output: empty string
     explanation:
          step 1: 112233444321 -> 112233321
          step 2: 112233321 -> 112221
          step 3: 112221 -> 111
          step 4: 111 -> ""

 Example 2:
     input: 11223344431
     output: 11221
     explanation:
          step 1: 11223344431 -> 11223331
          step 2: 112233321 -> 11221
