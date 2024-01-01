#include "orderbookentry.hpp"
#include <iostream>

OrderBookEntry::OrderBookEntry(double _price, double _amount, std::string _timestamp, std::string _product, OrderBookType _order_type)
    : price(_price), amount(_amount), timestamp(std::move(_timestamp)), product(std::move(_product)), order_type(_order_type) { }

OrderBookType OrderBookEntry::StringToOBT(std::string s) {
    if (s == "ask") {
        return OrderBookType::ask;
    }
    if (s == "bid") {
        return OrderBookType::bid;
    }

    return OrderBookType::unknown;
}

OrderBookType OrderBookEntry::CheckOrderType(OrderBookEntry obe) {
    if (obe.order_type == OrderBookType::ask)
        return OrderBookType::ask;
    if (obe.order_type == OrderBookType::bid)
        return OrderBookType::bid;

    return OrderBookType::unknown;
}

void OrderBookEntry::GetValue() const {
    // std::array<std::string, 2> enum_string = {"ask", "bid"};

    std::cout << "price: " << price << ";\t";
    std::cout << "amount: " << amount << ";\t";
    std::cout << "timestamp: " << timestamp << ";\t";
    std::cout << "product: " << product << '\n';
    // std::cout << "ordertype: \n" << order_type;
}
