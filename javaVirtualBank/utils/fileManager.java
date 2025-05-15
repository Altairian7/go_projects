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
