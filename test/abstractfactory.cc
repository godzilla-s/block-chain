#include <iostream>

using namespace std; 

class Widget
{
public:
    virtual void draw() = 0;
};

class LinuxButton : public Widget 
{
public:
    void draw() { cout << "LinuxButton\n"; }
};

class LinuxMenu : public Widget
{
public:
    void draw() { cout << "LinuxMenu\n"; }
};

class WindowButton : public Widget
{
public:
    void draw() { cout << "WindowButton\n"; }
};

class WindowMenu : public Widget
{
public:
    void draw() { cout << "WindowMenu\n"; }
};

class Factory
{
public:
    virtual Widget *create_button() = 0;
    virtual Widget *create_menu() = 0;
};

class LinuxFactory : public Factory
{
public:
    Widget *create_button() 
    {
        return new LinuxButton();
    }

    Widget *create_menu()
    {
        return new LinuxMenu();
    }
};

class WindowFactory : public Factory
{
public:
    Widget *create_button()
    {
        return new WindowButton;
    }

    Widget *create_menu()
    {
        return new WindowMenu;
    }
};

class Client
{
private:
    Factory *factory; 

public:
    Client(Factory *f)
    {
        factory = f;
    }

    void draw() 
    {
        Widget *w = factory->create_button();
        w->draw();
        display_window_one();
        display_window_two();
    }

    void display_window_one()
    {
        Widget *w[] = {
            factory->create_button(),
            factory->create_menu()
        };

        w[0]->draw();
        w[1]->draw();
    }

    void display_window_two()
    {
        Widget *w[] = {
            factory->create_button(),
            factory->create_menu()
        };

        w[0]->draw();
        w[1]->draw();
    }
};

int main()
{
    Factory *linux = new LinuxFactory;
    Factory *window = new WindowFactory;

    Client *c = new Client(linux);

    c->draw();
    delete c;
    c = new Client(window);
    c->draw();
    delete c;
}

// https://sourcemaking.com/design_patterns/builder/cpp/1