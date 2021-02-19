# Builder 

**Builder** is a creational design pattern, which allows constructing complex objects step by step. 

## Practical uses

The Builder is used when the designed product is complex and requires multiple steps to complete. In that case, several construction methods would be simpler than a single monstrous constructor. The potential problem with the multistage building process is tha a partially built and unstable product may be exposed to the client. The Builder pattern keeps the product private until it's fully built.  