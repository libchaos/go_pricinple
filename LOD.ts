class Teacher {
  command_gl_count(girlLeader: GirlsLeader): number {
    return girlLeader.count()
  }
}

class Girl {
  name = ''
  constructor(name: string) {
    this.name = name
  }

}

class GirlsLeader {
  public girls: Girl[] = []
  count(): number {
    return this.girls.length
  }
}

const t = new Teacher()

const g1 = new Girl('a')
const g2 = new Girl('b')

const gl = new GirlsLeader();
gl.girls.push(g1)
gl.girls.push(g2)


console.log(t.command_gl_count(gl))


