#include <iostream>
#include <string.h>
#include <stdio.h>

using namespace std;

enum PersitenceType
{
    File, Queue, Pathway
};

struct PersitenceAttribute
{
    PersitenceType type;
    char value[30];
};

class DistributeWorkPack
{
private:
    char    desc[200];
    char    temp[100];
public:
    DistributeWorkPack(char *type)
    {
        sprintf(desc, "Distribute work package for: %s", type);
    }

    void setFile(char *f, char *v)
    {
        sprintf(temp, "\n File(%s): %s", f, v);
        strcat(desc, temp);
    }

    void setQueue(char *q, char *v)
    {
        sprintf(temp, "\n Queue(%s): %s", q, v);
        strcat(desc, temp);
    }

    void setPathway(char *p, char *v)
    {
        sprintf(temp, "\n Pathway(%s): %s", p, v);
        strcat(desc, temp);
    }

    const char *getState()
    {
        return desc;
    }
};

class Builder
{
protected:
    DistributeWorkPack *result;
public:
    virtual void configureFile(char *) = 0;
    virtual void configureQueue(char *) = 0;
    virtual void configurePathway(char *) = 0;

    DistributeWorkPack *getResult()
    {
        return result;
    }
};

class UnixBuilder : public Builder
{
public:
    UnixBuilder()
    {
        result = new DistributeWorkPack("Unix");
    }

    void configureFile(char *name)
    {
        result->setFile("flatPath", name);
    }

    void configureQueue(char *queue)
    {
        result->setQueue("FIFO", queue);
    }

    void configurePathway(char *type)
    {
        result->setPathway("thread", type);
    }
};

class VmsBuilder : public Builder
{
public:
    VmsBuilder()
    {
        result = new DistributeWorkPack("Vms");
    }

    void configureFile(char *name)
    {
        result->setFile("ISAM", name);
    }

    void configureQueue(char *queue)
    {
        result->setQueue("priority", queue);
    }

    void configurePathway(char *type)
    {
        result->setPathway("LWP", type);
    }
};

class Reader
{
public:
    void setBuidler(Builder *b)
    {
        builder = b;
    }

    void construct(PersitenceAttribute[], int);

private:
    Builder *builder;
};

void Reader::construct(PersitenceAttribute list[], int num)
{
    for(int i=0; i<num; i++)
    {
        if (list[i].type == File)
            builder->configureFile(list[i].value);
        else if (list[i].type == Queue)
            builder->configureQueue(list[i].value);
        else if (list[i].type == Pathway)
            builder->configurePathway(list[i].value);
    }
}

const int NUM_ENTRIES = 6;
PersitenceAttribute input[NUM_ENTRIES] = 
{
    {File, "state.dat"},
    {File, "config.sys"},
    {Queue, "compute"},
    {Queue, "log"},
    {Pathway, "authentication"},
    {Pathway, "error processing"}
};

int main()
{
    UnixBuilder unixBuilder;
    VmsBuilder vmsBuilder;
    Reader  reader;

    reader.setBuidler(&unixBuilder);
    reader.construct(input, NUM_ENTRIES);
    cout << unixBuilder.getResult()->getState() << endl;
    
    reader.setBuidler(&vmsBuilder);
    reader.construct(input, NUM_ENTRIES);
    cout << vmsBuilder.getResult()->getState() << endl;
}