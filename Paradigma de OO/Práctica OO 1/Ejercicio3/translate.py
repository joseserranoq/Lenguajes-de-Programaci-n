
class Translate:
    def __init__(self):
        self.validate = ["-", "+", "*", "/"]
        self.operation = " "
        self.lOperators = []
        self.lNumbers = []
        self.result = 0
    def cal(self):
        if len(self.lOperators)==0:
            return 0
        else:
            if self.lOperators[0] == "+":
                self.result += int(self.lNumbers[0]) + int(self.lNumbers[1])
            elif self.lOperators[0] == "-":
                self.result += int(self.lNumbers[0]) - int(self.lNumbers[1])
            elif self.lOperators[0] == "*":
                self.result += int(self.lNumbers[0]) * int(self.lNumbers[1])
            elif self.lOperators[0] == "/":
                self.result += int(self.lNumbers[0]) / int(self.lNumbers[1])
            self.lOperators.pop(0)
            self.lNumbers.pop(0)
            self.lNumbers.pop(0)
            self.lNumbers.insert(0, self.result)
            return self.result + self.cal()

    def translate(self, operation):
        proof =operation.split(" ")
        try:
            for i in range(len(proof)):
                if proof[i] in self.validate:
                    self.lOperators.append(proof[i])
                else:
                    self.lNumbers.append(proof[i])
            print(self.cal())
        except Exception :
            print("Datos incorrectos", operation, Exception)
    
if __name__ == "__main__":
    translate = Translate()
    translate.translate("2 + 3 * 4 / 3 - 6 / 3 * 3 + 8")
            
