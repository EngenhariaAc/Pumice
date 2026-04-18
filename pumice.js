class Pumice {
  constructor(density, color) {
    this.density = density;
    this.color = color;
  }

  floatTest() {
    if (this.density < 1) {
      console.log("This pumice floats on water!");
    } else {
      console.log("This pumice sinks.");
    }
  }
}

module.exports = Pumice;
