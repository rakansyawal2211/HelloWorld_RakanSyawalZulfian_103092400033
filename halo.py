import time


beat = 0.7


lyrics = [
    "Did we ever know?",
    "Did we ever know?",
    "Did we ever know how strong we could be?",
    "",
    "Like lightning crashing on the sea",
    "",
    "Forget about where we are and let go",
    "We're so close"
]


for line in lyrics:
    print(line)
    time.sleep(beat * 4)  # jeda per baris
