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
