class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        mapper = {
            "1": [""],
            "2": ["a", "b", "c"],
            "3": ["d", "e", "f"],
            "4": ["g", "h", "i"],
            "5": ["j", "k", "l"],
            "6": ["m", "n", "o"],
            "7": ["p", "q", "r", "s"],
            "8": ["t", "u", "v"],
            "9": ["w", "x", "y", "z"]
        }
        output = []
        if len(digits): self.rec("", digits, mapper, 0, output)
        
        return output
    
    
    def rec(self, combi: str, digits: str, mapper: dict, k: int,
            output: List[str]) -> List[str]:
        
        if len(digits) == k:
            output.append(combi)
            return
        
        current_digit = digits[k]
        for i in range(len(mapper[current_digit])):
            combi+=mapper[current_digit][i]
            self.rec(combi, digits, mapper, k+1, output)
            combi=combi[:-1]
            
