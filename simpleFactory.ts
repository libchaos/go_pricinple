interface IGril {
  weight(): void
}

class FatGirl {
  weight() {
    console.log('fat');
  }
}

class ThinGirl {
  weight() {
    console.log('thin');
  }
}

enum Weight {
  thin,
  fat
}

class GirlFactory {
  createGirl(like: Weight): IGril{
    switch (like) {
      case Weight.thin:
        return new FatGirl()
      case Weight.fat:
        return new ThinGirl()
    }
  }
}

const gf = new GirlFactory()

const thinGirl = gf.createGirl(Weight.thin)
thinGirl.weight()

