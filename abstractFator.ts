interface Factory {
  newTV(): Televison
  newRefrigerator(): Refrigerator
}

interface Televison {
  playing(): void
}

interface Refrigerator {
  cooling(): void
}


class TCLFactory implements Factory {
  newTV(): Televison {
    return new TCLTV()
  }
  newRefrigerator(): Refrigerator {
    return new TCLRef()
  }
}

class TCLTV implements Televison{
  playing() {
    console.log("TCL playing");
  }
}

class TCLRef {
  cooling() {
    console.log("cooling");
  }
}

class MediaFactory {
  newTV(): Televison {
    return new MediaTV()
  }
  newRefrigerator(): Refrigerator {
    return new MediaRef()
  }
}

class MediaTV implements Televison {
  playing() {
    console.log("Media playing");
  }
}

class MediaRef implements Refrigerator{ 
  cooling() {
    console.log("Media cooling");
  }
}

