/*
 * Step-by-step implementation of the Java Virtual Bank Project
 * Covers complete Java Syllabus Topics
 */

// STEP 1: Create Main.java

package main;

import ui.BankUI;

public class Main {
    public static void main(String[] args) {
        new BankUI();
    }
}


// STEP 2: Create models/User.java

package models;

import java.io.Serializable;

public class User implements Serializable {
    private String username;
    private String password;
    private String fullName;

    public User(String username, String password, String fullName) {
        this.username = username;
        this.password = password;
        this.fullName = fullName;
    }

    public String getUsername() { return username; }
    public String getPassword() { return password; }
    public String getFullName() { return fullName; }
}


// STEP 3: Create models/Account.java

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


// STEP 4: Create models/Transaction.java

package models;

import java.io.Serializable;
import java.util.Date;

public class Transaction implements Serializable {
    private Date date;
    private String type;
    private double amount;

    public Transaction(String type, double amount) {
        this.date = new Date();
        this.type = type;
        this.amount = amount;
    }

    public String toString() {
        return date + " - " + type + ": Rs." + amount;
    }
}


// STEP 5: Create utils/FileManager.java

package utils;

import models.*;

import java.io.*;
import java.util.HashMap;

public class FileManager {
    private static final String FILE = "data/users.dat";

    public static HashMap<String, User> readUsers() {
        try (ObjectInputStream ois = new ObjectInputStream(new FileInputStream(FILE))) {
            return (HashMap<String, User>) ois.readObject();
        } catch (Exception e) {
            return new HashMap<>();
        }
    }

    public static void writeUsers(HashMap<String, User> users) {
        try (ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(FILE))) {
            oos.writeObject(users);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}


// STEP 6: Create services/BankService.java

package services;

import models.*;
import java.util.*;

public class BankService {
    private User currentUser;
    private Account account;
    private List<Transaction> transactions;

    public BankService(User user) {
        this.currentUser = user;
        this.account = new Account();
        this.transactions = new ArrayList<>();
    }

    public void deposit(double amt) {
        account.deposit(amt);
        transactions.add(new Transaction("Deposit", amt));
    }

    public boolean withdraw(double amt) {
        boolean result = account.withdraw(amt);
        if (result) {
            transactions.add(new Transaction("Withdraw", amt));
        }
        return result;
    }

    public double getBalance() {
        return account.getBalance();
    }

    public List<Transaction> getTransactions() {
        return transactions;
    }
}


// STEP 7: Create ui/BankUI.java

package ui;

import models.User;
import services.BankService;
import utils.FileManager;

import java.awt.*;
import java.awt.event.*;
import java.util.HashMap;

public class BankUI extends Frame {
    private TextField usernameTF, passwordTF, amountTF;
    private Label statusLabel, balanceLabel;
    private BankService service;
    private HashMap<String, User> users;

    public BankUI() {
        users = FileManager.readUsers();
        setTitle("Java Virtual Bank");
        setSize(400, 300);
        setLayout(new GridLayout(6, 2));

        add(new Label("Username:"));
        usernameTF = new TextField();
        add(usernameTF);

        add(new Label("Password:"));
        passwordTF = new TextField();
        passwordTF.setEchoChar('*');
        add(passwordTF);

        Button loginBtn = new Button("Login");
        Button registerBtn = new Button("Register");
        add(loginBtn);
        add(registerBtn);

        loginBtn.addActionListener(e -> login());
        registerBtn.addActionListener(e -> register());

        add(new Label("Amount:"));
        amountTF = new TextField();
        add(amountTF);

        Button depositBtn = new Button("Deposit");
        Button withdrawBtn = new Button("Withdraw");
        add(depositBtn);
        add(withdrawBtn);

        depositBtn.addActionListener(e -> deposit());
        withdrawBtn.addActionListener(e -> withdraw());

        balanceLabel = new Label("Balance: Rs.0");
        add(balanceLabel);

        statusLabel = new Label();
        add(statusLabel);

        addWindowListener(new WindowAdapter() {
            public void windowClosing(WindowEvent e) {
                FileManager.writeUsers(users);
                System.exit(0);
            }
        });

        setVisible(true);
    }

    private void login() {
        String user = usernameTF.getText();
        String pass = passwordTF.getText();
        if (users.containsKey(user) && users.get(user).getPassword().equals(pass)) {
            statusLabel.setText("Login successful!");
            service = new BankService(users.get(user));
        } else {
            statusLabel.setText("Invalid credentials");
        }
    }

    private void register() {
        String user = usernameTF.getText();
        String pass = passwordTF.getText();
        if (!users.containsKey(user)) {
            User newUser = new User(user, pass, user);
            users.put(user, newUser);
            statusLabel.setText("Registered successfully");
        } else {
            statusLabel.setText("User already exists");
        }
    }

    private void deposit() {
        try {
            double amt = Double.parseDouble(amountTF.getText());
            service.deposit(amt);
            balanceLabel.setText("Balance: Rs." + service.getBalance());
        } catch (Exception e) {
            statusLabel.setText("Invalid amount");
        }
    }

    private void withdraw() {
        try {
            double amt = Double.parseDouble(amountTF.getText());
            if (service.withdraw(amt)) {
                balanceLabel.setText("Balance: Rs." + service.getBalance());
            } else {
                statusLabel.setText("Insufficient Balance");
            }
        } catch (Exception e) {
            statusLabel.setText("Invalid amount");
        }
    }
}
