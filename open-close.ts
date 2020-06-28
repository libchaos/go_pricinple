interface ICar {
  GetName(): string;
  GetPrice(): number;
}

class BenzCar implements ICar {
  public name: string = "";
  public price: number = 0;
  constructor(name: string, price: number) {
    this.name = name;
    this.price = price;
  }

  GetName() {
    return this.name;
  }

  GetPrice(): number {
    return this.price;
  }
}

class FianceBenzCar extends BenzCar implements ICar {
  constructor(name: string, price: number) {
    super(name, price);
  }
  GetPrice(): number {
    let price = this.price;
    if (price >= 100) {
      this.price = price + price * 5 / 100;
    } else if (price >= 50) {
      this.price = price + price * 2 / 100;
    } else {
      this.price = price;
    }
    return this.price;
  }
}

function main() {
  let list: Array<ICar> = [
    new BenzCar("amg", 1000),
    new FianceBenzCar("omg", 2131),
  ];
  list.forEach((item) => {
    console.log(`${item.GetName()}  ${item.GetPrice()}`);
  });
}

main();
