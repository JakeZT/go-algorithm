class MedianFinder {
	store: number[]
	constructor() {
		this.store = []
	}

	addNum(num: number): void {
		if (!this.store.length) {
			this.store.push(num)
		} else {
			let i = this.store.length - 1
			let pos = this.store.length
			while (i >= 0 && num < this.store[i]) {
				if (i < 0) {
					break
				}
				pos--
				i--
			}
			this.store.splice(pos, 0, num)
		}
	}

	findMedian(): number {
		let len = this.store.length
		let midIndex = Math.floor(this.store.length / 2)
		if (len % 2 == 0) {
			return (this.store[midIndex] + this.store[midIndex - 1]) * 0.5
		} else {
			return this.store[midIndex]
		}
	}
}
