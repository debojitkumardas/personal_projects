SRC_DIR = src/
HDR_DIR = hdr/
BIN_DIR = bin/

CXX = g++
DEPFLAGS = -MP -MD
CXXFLAGS = -Wall -Weffc++ -Wextra -Wconversion -Wsign-conversion -Werror -std=c++20 -I$(HDR_DIR) $(DEPFLAGS)
LIBS = -lpthread

SRC = $(wildcard $(SRC_DIR)*.cpp)
HDR = $(wildcard $(HDR_DIR)*.hpp)
BIN = $(BIN_DIR)main

OBJ_O = $(patsubst %.cpp, $(BIN_DIR)%.o, $(notdir $(SRC)))
HDR_D = $(patsubst %.hpp, $(BIN_DIR)%.d, $(notdir $(HDR)))

all: $(BIN)

debug: CXXFLAGS += -DDEBUG -ggdb
debug: $(BIN)

release: CXXFLAGS += -O2 -DNDEBUG
release: $(BIN)

$(BIN): $(OBJ_O)
	$(CXX) $(CXXFLAGS) -o $@ $^ $(LIBS)

$(BIN_DIR)%.o: $(SRC_DIR)%.cpp $(HDR)
	$(CXX) $(CXXFLAGS) -c -o $@ $< $(LIBS)

clean:
	rm -rf $(OBJ_O) $(HDR_D) $(BIN)

clean_o:
	rm -rf .cache compile_commands.json
