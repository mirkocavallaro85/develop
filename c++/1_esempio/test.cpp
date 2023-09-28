/*
  Il primo programma di esempio:
  HelloWorld!
*/
#include <iostream>
// Definizione della classe Greater
class Greater
{
public:
    void sayHello()
    {
        std::cout << "Hello World!" << std::endl;
    }
};

// Funzione main: il punto di inizio per l'esecuzione del programma
int main()
{
    Greater g;
    g.sayHello();
    return 0;
}