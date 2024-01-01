SRC_DIR = src/
HDR_DIR = hdr/
BIN_DIR = bin/

CXX = g++
DEPFLAGS = -MP -MD
CXXFLAGS = -Wall -Weffc++ -Wextra -Wconversion -Wsign-conversion -Werror -pedantic -std=c++20 -I$(HDR_DIR) $(DEPFLAGS)

SRC = $(wildcard $(SRC_DIR)*.cpp)
HDR = $(wildcard $(HDR_DIR)*.hpp)
BIN = $(BIN_DIR)main

HDR_D = $(patsubst %.cpp, $(BIN_DIR)%.d, $(notdir $(SRC)))
OBJ_O = $(patsubst %.cpp, $(BIN_DIR)%.o, $(notdir $(SRC)))

all: $(BIN)

debug: CXXFLAGS += -DDEBUG -ggdb
debug: $(BIN)

release: CXXFLAGS += -O2 -DNDEBUG
release: $(BIN)

$(BIN): $(OBJ_O)
	$(CXX) $(CXXFLAGS) -o $@ $^

$(BIN_DIR)%.o: $(SRC_DIR)%.cpp $(HDR)
	$(CXX) $(CXXFLAGS) -c -o $@ $<

run_with_valgrind:
	valgrind --enable-debuginfod=no --leak-check=full --show-leak-kinds=all --errors-for-leak-kinds=all --track-origins=yes -s --verbose --log-file=valile=valgrind_out.txt ./bin/main

clean:
	rm -rf $(HDR_D) $(OBJ_O) $(BIN)

clean_o:
	rm -rf .cache compile_commands.json
