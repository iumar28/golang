#include<bits/stdc++.h>
using namespace std;

class Person{
    public:
        virtual void print_person()=0;    
};
class student: public Person{
    public:
    void print_person(){
            cout<<"This is the student class\n";
    }
};
class faculty: public Person{
    public:
    void print_peron(){
        cout<<"This is the faculty class\n";
    }
};
class college{
    private:
    Person *guest;
    public:
    college(int type ){
        if(type==1)
        {guest=new student();}
        else if(type==2)
        {guest=new faculty();}
        else
        {guest=NULL;}
    }
    ~college(){
        if(guest){
            delete[] guest;
            guest=nullptr;
        }
    }
    college* getperson()
    {return guest;}
};

int main(){
    college *col=new college(1);
    Person *check=col->getperson();
    check->print_person();
    return 0;


}
