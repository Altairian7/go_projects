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