class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        
        for letter in s:
            if stack and self.end_paren(letter):
                if self.is_correctly_closed(stack[-1], letter):
                    stack.pop()
                else:
                    break
            else:
                stack.append(letter)
        
        return not len(stack)
    
    def is_correctly_closed(self, begin, end):
        return begin == '(' and end == ')' or begin == '{' and end == '}' or begin == '[' and end == ']'
    
    def end_paren(self, s):
        return s == ')' or s == '}' or s == ']'
