# string-reduce

## Given a string representing a list of single characters, remove any group of the same characters if there are at least N consecutive occurrences. Repeat the same treatment to the string after removal until there is no more group of same character with N or more consecutive occurrences

```bash
Example 1:
     ./binaries/string-reduce -consecutive 3 -input 112233444321
     Input string 112233444321 has been reduced to 
     +---------------------------------+
     input: 112233444321
     output: empty string
     explanation:
          step 1: 112233444321 -> 112233321
          step 2: 112233321 -> 112221
          step 3: 112221 -> 111
          step 4: 111 -> ""

 Example 2:
     ./binaries/string-reduce -consecutive 3 -input 11223344431
     Input string 11223344431 has been reduced to 11221
     +---------------------------------+
     input: 11223344431
     output: 11221
     explanation:
          step 1: 11223344431 -> 11223331
          step 2: 112233321 -> 11221
```
