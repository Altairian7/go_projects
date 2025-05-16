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
