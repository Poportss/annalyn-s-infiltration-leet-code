# Annalyn's Infiltration

This repository contains the solution for the **"Annalyn's Infiltration"** challenge.

## 📜 Description
Annalyn and her team are planning to rescue a prisoner. Several conditions determine whether certain actions can be executed. This code implements functions to evaluate these conditions based on the status of different characters.

## 🚀 Features

- **`CanFastAttack(knightIsAwake bool) bool`**  
  🛡️ **Description:** Returns `true` if a fast attack can be executed, meaning the knight is asleep.
  
- **`CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool`**  
  🔍 **Description:** Returns `true` if at least one of the characters is awake.
  
- **`CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool`**  
  📢 **Description:** Returns `true` if the prisoner is awake and the archer is asleep.
  
- **`CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool`**  
  🔓 **Description:** Returns `true` if the prisoner is awake while the knight and archer are asleep, or if Annalyn's pet dog is present and the archer is asleep.

## 🛠️ How to Run

1. Clone the repository:
   ```sh
   git clone https://github.com/Poportss/annalyn-infiltration.git
   cd annalyn-infiltration
   ```
2. Compile and run the code:
   ```sh
   go run main.go
   ```

## 🖥️ Example Output

```sh
Can Fast Attack: true
Can Spy: true
Can Signal Prisoner: false
Can Free Prisoner: true
```

## 🏗️ Technologies Used
- [Go (Golang)](https://golang.org/)

## ✨ Author
Rafael Portela

