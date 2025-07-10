class Animal{
    constructor(name,voice){
        this.name=name;
        this.voice=voice;
    }

    speak(){
        console.log(this.voice);
    }
}

class Dog extends Animal{
    
    bark(){
        console.log("Woof!");
    }
}

const Adog= new Dog("buggy","bark");
Adog.speak();
Adog.bark();
