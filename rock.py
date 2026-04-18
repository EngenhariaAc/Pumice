class Rock:
    def __init__(self, name, porosity):
        self.name = name
        self.porosity = porosity

    def describe(self):
        return f"{self.name} has a porosity of {self.porosity}%."
