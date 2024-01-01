#pragma once

#include <string>

enum class OrderBookType {bid, ask, unknown};

class OrderBookEntry {
private:
    double price;
    double amount;
    std::string timestamp;
    std::string product;
    OrderBookType order_type;

public:
    OrderBookEntry(double _price, double _amount, std::string _timestamp, std::string _product, OrderBookType _order_type);
    static OrderBookType StringToOBT(std::string);
    OrderBookType CheckOrderType(OrderBookEntry);
    void GetValue() const;
};
