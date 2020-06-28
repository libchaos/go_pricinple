interface AbstractAniaml {
  dance():void
}

class Rabiit implements AbstractAniaml {
  dance() {
    console.log("兔子跳舞")
  }
}


class Dog implements AbstractAniaml {
  dance() {
    console.log("狗子跳舞")
  }
}


class Person {
  constructor(public animal: AbstractAniaml) {}
  walkAnimal() {
    console.log("人开始遛动物")
    this.animal.dance()
  }
}


const p = new Person(new Dog())
p.walkAnimal()