module Finance
  def self.rate(nper : Int32, pmt : Float64, pv : Float64, fv : Float64, loan_type = 0, guess = 0.1)
    tolerancy = 1e-6
    close = false
    next_guess = 0

    while !close
      temp = newton_iter(guess, nper, pmt, pv, fv, loan_type)
      next_guess = (guess - temp).round(20)
      diff = (next_guess - guess).abs
      close = diff < tolerancy
      guess = next_guess
    end

    (next_guess * 12 * 100).round(2)
  end

  private def self.newton_iter(r, n, p, x, y, w)
    t1 = (r + 1)**n
    t2 = (r + 1)**(n - 1)
    ((y + t1*x + p*(t1 - 1)*(r*w + 1)/r) / (n*t2*x - p*(t1 - 1)*(r*w + 1)/(r**2) + n*p*t2*(r*w + 1)/r + p*(t1 - 1)*w/r))
  end
end