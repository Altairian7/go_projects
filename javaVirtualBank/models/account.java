package models;

import java.io.Serializable;

public class Account implements Serializable {
    private double balance;

    public Account() {
        this.balance = 0.0;
    }

    public synchronized void deposit(double amount) {
        balance += amount;
    }

    public synchronized boolean withdraw(double amount) {
        if (balance >= amount) {
            balance -= amount;
            return true;
        }
        return false;
    }

    public double getBalance() {
        return balance;
    }
}